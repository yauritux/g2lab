package com.yauritux.service.query.impl;

import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.yauritux.model.entity.Card;
import com.yauritux.model.entity.CardTransaction;
import com.yauritux.repository.CardTransactionRepository;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@Service
public class CardTransactionQueryServiceImpl implements com.yauritux.service.query.CardTransactionQueryService {

	private CardTransactionRepository cardTransactionRepository;
	
	@Autowired
	CardTransactionQueryServiceImpl(CardTransactionRepository cardTransactionRepository) {
		this.cardTransactionRepository = cardTransactionRepository;
	}
	
	@Override
	public Optional<List<CardTransaction>> findByCard(Card card) {
		return Optional.ofNullable(cardTransactionRepository.findByCard(card).stream().collect(Collectors.toList()));
	}
}
