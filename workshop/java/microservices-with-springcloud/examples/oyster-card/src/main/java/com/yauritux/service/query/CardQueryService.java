package com.yauritux.service.query;

import java.util.Optional;

import com.yauritux.model.entity.Card;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
public interface CardQueryService {

	Optional<Card> findBySerialNo(final String serialNo);
}
