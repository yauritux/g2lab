package com.yauritux.service.query.impl;

import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.yauritux.model.entity.Card;
import com.yauritux.repository.CardRepository;
import com.yauritux.service.query.CardQueryService;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@Service
public class CardQueryServiceImpl implements CardQueryService {
	
	private CardRepository cardRepository;
	
	@Autowired
	CardQueryServiceImpl(CardRepository cardRepository) {
		this.cardRepository = cardRepository;
	}

	@Override
	public Optional<Card> findBySerialNo(String serialNo) {
		return Optional.ofNullable(cardRepository.findBySerialNo(serialNo));
	}

}
