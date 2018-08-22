package com.yauritux.service.command.impl;

import java.math.BigDecimal;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.yauritux.exception.CardException;
import com.yauritux.model.builder.CardBuilder;
import com.yauritux.model.constant.CardOperationType;
import com.yauritux.model.entity.Card;
import com.yauritux.repository.CardRepository;
import com.yauritux.service.command.CardCommandService;
import com.yauritux.service.validator.CardValidator;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@Service
public class CardCommandServiceImpl implements CardCommandService {
	
	private CardRepository cardRepository;
	private CardValidator cardValidator;
	
	@Autowired
	CardCommandServiceImpl(CardRepository cardRepository, CardValidator cardValidator) {
		this.cardRepository = cardRepository;
		this.cardValidator = cardValidator;
	}
	
	@Override
	public Card initializeCard(String owner, BigDecimal initialBalance) throws CardException {
		CardException cardException = cardValidator.validateOwner(owner);
		if (cardException != null) {
			throw cardException;
		}
		
		cardException = cardValidator.validateAmount(initialBalance, CardOperationType.CARD_INITIALIZATION);
		if (cardException != null) {
			throw cardException;
		}
		
		return new CardBuilder(owner).setBalance(initialBalance).build();
	}	

	@Override
	public BigDecimal deduct(String serialNo, BigDecimal amount) throws CardException {
		CardException cardException = cardValidator.validateSerialNumber(serialNo);
		if (cardException != null) {
			throw cardException; 
		}
		cardException = cardValidator.validateAmount(amount, CardOperationType.DEDUCT_BALANCE);
		if (cardException != null) {
			throw cardException;
		}
		Card card = cardRepository.findBySerialNo(serialNo);
		card.setBalance(card.getBalance().subtract(amount));
		cardRepository.save(card);
		return card.getBalance();
	}

	@Override
	public BigDecimal add(String serialNo, BigDecimal amount) throws CardException {
		CardException cardException = cardValidator.validateSerialNumber(serialNo);
		if (cardException != null) {
			throw cardException;
		}
		cardException = cardValidator.validateAmount(amount, CardOperationType.ADD_BALANCE);
		if (cardException != null) {
			throw cardException;
		}
		Card card = cardRepository.findBySerialNo(serialNo);
		card.setBalance(card.getBalance().add(amount));
		cardRepository.save(card);
		return card.getBalance();
	}
}
